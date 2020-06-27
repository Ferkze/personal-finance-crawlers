const fs = require('fs');

console.info('Running shares script')

let orders = JSON.parse(fs.readFileSync('orders.json'));

const shares = orders.filter(o => o.assetType == 'Ações')
	.sort((a,b) => new Date(b.date) - new Date(a.date))
	.map(o => {
		o.date = o.date.substring(0, 10)
		return o
	})

const dateKeys = [...new Set(shares.map(o => o.date))]
let results = dateKeys.map(k => {
	const trades = shares.filter(o => o.date == k)
	const positions = {}
	for (const trade of trades) {
		if (!positions[trade.asset]) {
			positions[trade.asset] = {
				totalVolume: 0,
				buyingVolume: 0,
				sellingVolume: 0,
				buyQnt: 0,
				sellingQnt: 0,
			}
		}
		if (trade.qnt > 0) {
			positions[trade.asset].buyingVolume += trade.price * trade.qnt
			positions[trade.asset].buyQnt += Math.abs(trade.qnt)
		} else {
			positions[trade.asset].sellingVolume += trade.price * Math.abs(trade.qnt)
			positions[trade.asset].sellingQnt += Math.abs(trade.qnt)
		}
		positions[trade.asset].totalVolume += Math.abs(trade.qnt) * trade.price
		positions[trade.asset].date = trade.date
	}

	const result = calculatePositionsResult(positions)
	result.date = k
	result.month = k.substring(0,7)
	
	return { [k]: result }
})

results = groupResultsByPeriod(results, 'month')

printResults(results)

function calculatePositionsResult(positions = {'ASSET': {totalVolume: 0, buyingVolume: 0, sellingVolume: 0, buyQnt: 0, sellingQnt: 0}}) {
	const results = {
		total: 0,
		costs: 0,
		irrf: 0
	}
	for (const position of Object.values(positions)) {
		if (position.sellingQnt == 0 || position.buyQnt == 0) {
			continue
		}
		if (position.buyQnt > position.sellingQnt && position.sellingQnt > 0) {
			const excedent = position.buyQnt - position.sellingQnt
			const buyVolDT = position.buyingVolume - ((position.buyingVolume/position.buyQnt)*excedent)
			const res = position.sellingVolume - buyVolDT
			results.total += res
		} else if (position.sellingQnt > position.buyQnt && position.buyQnt > 0) {
			const excedent = position.sellingQnt - position.buyQnt
			const shortVolDT = position.sellingVolume - ((position.sellingVolume/position.sellingQnt)*excedent)
			const res = shortVolDT - position.buyingVolume
			results.total += res
		} else {
			const res = position.sellingVolume - position.buyingVolume
			results.total += res
		}
	}
	results.costs = results.total * 0.0003084
	if (results.total - results.costs > 0) {
		results.irrf = (results.total - results.costs) * 0.01
	}

	return results
}

function groupResultsByPeriod(results, period='month') {
	return results.reduce((acc, cur) => {
		const key = Object.keys(cur)[0]
		let periodKey
		switch(period) {
			case 'month':
				periodKey = cur[key].month
				break
			case 'year':
				periodKey = cur[key].month.substring(0,4)
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

function printResults(results) {
	Object.keys(results).forEach(k => {
		const r = results[k]
		console.log(`resultado do mês ${k}: R$ ${(r.total - r.costs).toFixed(2).padStart(8, ' ')}, irrf: R$ ${r.irrf.toFixed(2).padStart(8, ' ')}`)
	})

}