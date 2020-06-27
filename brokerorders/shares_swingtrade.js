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

	const result = calculateSwingtradesResult(positions)
	result.date = k
	result.month = k.substring(0,7)
	
	return { [k]: result }
})

results = groupResultsByPeriod(results, 'month')

printResults(results)


function calculateSwingtradesResult(positions = {'ASSET': {totalVolume: 0, totalBuyingVolume: 0, totalSellingVolume: 0, buyQnt: 0, sellingQnt: 0}}) {
	const results = {
		total: 0,
		sold: 0,
		costs: 0,
		irrf: 0
	}
	for (const position of Object.values(positions)) {
		if (position.operation == 'daytrade') {
			continue
		}
		// TODO: criar objeto que guarda o ativo e suas posições, total da posicao, quantidade e etc
		// TODO: calcular o valor total da posicao comprada ou vendida trade após trade, e todo o custo da negociação.
		// if (position.buyQnt > position.sellingQnt && position.sellingQnt > 0) {
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
	results.costs = results.total * 0.0003084
	if (results.total - results.costs > 0) {
		results.irrf = (results.total - results.costs) * 0.01
	}

	return results
}
