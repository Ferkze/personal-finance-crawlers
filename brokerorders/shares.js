const fs = require('fs');
const {
	filterOrdersByType,
	filterOrdersByDate,
	sortOrdersByDate
} = require('./utils')

console.info('Running shares script')

/**
 * @typedef {Object} Order
 * @property {string} asset
 * @property {string} assetType
 * @property {number} qnt
 * @property {number} price
 * @property {string} date
 */

/**
 * @type {Order[]}
 */
const ordersExtracted = JSON.parse(fs.readFileSync('orders.json'));

const orders = sortOrdersByDate(filterOrdersByType(ordersExtracted, 'Ações'))
	.map(o => ({ ...o, date: o.date.substring(0, 10) }))

const results = groupResultsByPeriod(generateResults(orders), 'month')

Object.keys(results).forEach(k => {
	const r = results[k]
	const format = val => val.toFixed(2).padStart(8, ' ')
	console.log(`resultado do mês ${k}: R$ ${format(r.total - r.costs)}, irrf: R$ ${format(r.irrf)}`)
})

function generateResults(orders) {
	const dateKeys = [...new Set(orders.map(o => o.date))]
	dateKeys.map(k => {
		const trades = filterOrdersByDate(orders, date)
		const positions = {}
		for (const trade of trades) {
			if (!positions[trade.asset]) {
				positions[trade.asset] = {
					totalVolume: 0,
					totalBuyingVolume: 0,
					totalSellingVolume: 0,
					sellingVolume: 0,
					buyQnt: 0,
					sellingQnt: 0,
					operation: ''
				}
			}
			if (trade.qnt > 0) {
				positions[trade.asset].totalBuyingVolume += trade.price * trade.qnt
				positions[trade.asset].buyQnt += Math.abs(trade.qnt)
			} else {
				positions[trade.asset].totalSellingVolume += trade.price * Math.abs(trade.qnt)
				positions[trade.asset].sellingQnt += Math.abs(trade.qnt)
			}
			if (positions[trade.asset].buyQnt == positions[trade.asset].sellingQnt) {
				positions[trade.asset].operation = 'daytrade'
			} else if (positions[trade.asset].buyQnt != 0 && positions[trade.asset].sellingQnt != 0) {
				positions[trade.asset].operation = 'mixed'
			} else {
				positions[trade.asset].operation = 'swingtrade'
			}
			if (positions[trade.asset].operation == 'swingtrade' && positions[trade.asset].sellingQnt > 0) {
				positions[trade.asset].sellingVolume += trade.price * Math.abs(trade.qnt)
			}
			positions[trade.asset].totalVolume += Math.abs(trade.qnt) * trade.price
			positions[trade.asset].date = trade.date
		}

		const result = calculatePositionsResult(positions)
		result.date = k
		result.month = k.substring(0, 7)

		return { [k]: result }
	})
}

function calculatePositionsResult(positions = { 'ASSET': { totalVolume: 0, totalBuyingVolume: 0, totalSellingVolume: 0, buyQnt: 0, sellingQnt: 0 } }) {
	const results = {
		total: 0,
		sold: 0,
		costs: 0,
		irrf: 0
	}
	for (const position of Object.values(positions)) {
		if (position.sellingQnt == 0 || position.buyQnt == 0) {
			continue
		}
		if (position.buyQnt > position.sellingQnt && position.sellingQnt > 0) {
			const excedent = position.buyQnt - position.sellingQnt
			const buyVolDT = position.totalBuyingVolume - ((position.totalBuyingVolume / position.buyQnt) * excedent)
			const res = position.totalSellingVolume - buyVolDT
			results.total += res
		} else if (position.sellingQnt > position.buyQnt && position.buyQnt > 0) {
			const excedent = position.sellingQnt - position.buyQnt
			const shortVolDT = position.totalSellingVolume - ((position.totalSellingVolume / position.sellingQnt) * excedent)
			const res = shortVolDT - position.totalBuyingVolume
			results.total += res
		} else {
			const res = position.totalSellingVolume - position.totalBuyingVolume
			results.total += res
		}
	}
	results.costs = results.total * 0.0003084
	if (results.total - results.costs > 0) {
		results.irrf = (results.total - results.costs) * 0.01
	}

	return results
}

function groupResultsByPeriod(results, period = 'month') {
	return results.reduce((acc, cur) => {
		const key = Object.keys(cur)[0]
		let periodKey
		switch (period) {
			case 'month':
				periodKey = cur[key].month
				break
			case 'year':
				periodKey = cur[key].month.substring(0, 4)
				break
		}
		if (!acc[periodKey]) {
			acc[periodKey] = {
				total: 0,
				irrf: 0,
				costs: 0,
				date: '',
				month: '',
			}
		}
		acc[periodKey].total += cur[key].total
		acc[periodKey].irrf += cur[key].irrf
		acc[periodKey].costs += cur[key].costs
		return acc
	}, {})
}
