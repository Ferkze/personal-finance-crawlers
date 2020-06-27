const fs = require('fs');

console.info('Running shares script')

let orders = JSON.parse(fs.readFileSync('orders.json'));

const shares = orders.filter(o => o.assetType == 'Ações')
	.sort((a,b) => new Date(b.date) - new Date(a.date))
	.map(o => {
		o.date = o.date.substring(0, 10)
		return o
	})

const swingtrades = {}
const dateKeys = [...new Set(shares.map(o => o.date))]
let results = dateKeys.map(k => {
	const trades = shares.filter(o => o.date == k)
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

	const result = calculateSwingtradesResult(positions, swingtrades)
	result.date = k
	result.month = k.substring(0,7)
	
	return { [k]: result }
})

results = groupResultsByPeriod(results, 'month')

printResults(results)


function calculateSwingtradesResult(positions = {'': {totalVolume: 0, totalBuyingVolume: 0, totalSellingVolume: 0, buyQnt: 0, sellingQnt: 0, operation: ''}}, swingtrades) {
	const results = {
		total: 0,
		volume: 0,
		sold: 0,
		costs: 0,
		irrf: 0
	}
	// let stcount = Object.values(positions).reduce((acc, cur) => cur.operation == 'swingtrade' ? ++acc : acc, 0)

	for (const positionKey of Object.keys(positions)) {
		const position = positions[positionKey]
		if (position.operation == 'daytrade') {
			continue
		}
		let st = swingtrades[positionKey]
		if (!st) {// nova posicao
			st = {
				total: 0,
				quant: 0
			}
		}
		if (position.operation == 'swingtrade') {
			if (position.buyQnt > 0) { // compra
				if (st.quant >= 0 && st.total >= 0) {// comprado
					st.total += position.totalBuyingVolume
					st.quant += position.buyQnt
				} else {// vendido
					let avgPriceShort = st.total/st.quant
					let avgPriceLong = position.totalBuyingVolume/position.buyQnt
					if (position.buyQnt <= st.quant) {
						results.total += (avgPriceShort - avgPriceLong)*position.buyQnt
						st.quant += position.buyQnt
						st.total += position.totalBuyingVolume
					} else {
						console.log('tratar 1')
					}
				}
				results.volume += position.totalBuyingVolume
			} else if(position.sellingQnt > 0) { // venda
				if (st.quant >= 0 && st.total >= 0) {// comprado
					let avgPriceShort = position.totalSellingVolume/position.sellingQnt
					let avgPriceLong = st.total/st.quant
					if (position.sellingQnt <= st.quant) {
						results.total += (avgPriceShort - avgPriceLong)*position.sellingQnt
						st.quant -= position.sellingQnt
						st.total -= position.totalSellingVolume
					} else {
						console.log('tratar 2')
					}
				} else {// vendido
					st.total -= position.totalSellingVolume
					st.quant += position.sellingQnt
				}
				results.volume += position.totalSellingVolume
				results.sold += position.totalSellingVolume
			}
		} else if (position.operation == 'mixed') {
			console.log('tratar 3')
			// if (position.buyQnt > 0 && position.sellingQnt > 0) {
			// 	const excedent = position.buyQnt - position.sellingQnt
			// 	const buyVolDT = position.totalBuyingVolume - ((position.totalBuyingVolume/position.buyQnt)*excedent)
			// 	const res = position.totalSellingVolume - buyVolDT
			// 	results.total += res
			// } else if (position.sellingQnt > position.buyQnt && position.buyQnt > 0) {
			// 	const excedent = position.sellingQnt - position.buyQnt
			// 	const shortVolDT = position.totalSellingVolume - ((position.totalSellingVolume/position.sellingQnt)*excedent)
			// 	const res = shortVolDT - position.totalBuyingVolume
			// 	results.total += res
			// } else {
			// 	const res = position.totalSellingVolume - position.totalBuyingVolume
			// 	results.total += res
			// }
		}
		swingtrades[positionKey] = st
		if (swingtrades[positionKey].quant == 0) {
			delete swingtrades[positionKey]
		}
	}
	results.costs = results.volume * 0.0003084
	if (results.sold > 0) {
		results.irrf = results.sold * 0.005
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
		console.log(`Período ${k}`)
		console.log(`Resultado Líquido: R$ ${(r.total - r.costs).toFixed(2).padStart(8, ' ')} ; Custos: R$ ${(r.costs).toFixed(2).padStart(8, ' ')} ; IRRF: R$ ${r.irrf.toFixed(2).padStart(8, ' ')}`)
	})
}