const fs = require('fs');

console.info('Running shares script')

let orders = JSON.parse(fs.readFileSync('orders.json'));

const shares = orders.filter(o => o.assetType == 'Ações')
	.sort((a,b) => new Date(a.date) - new Date(b.date))
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
		handleStockDivision(positionKey, position)
		if (position.operation == 'swingtrade') {
			consolidateResult(results, st, position, positionKey)
		} else if (position.operation == 'mixed') {
			let pos
			if (position.buyQnt > position.sellingQnt) {
				pos = {
					totalBuyingVolume: position.totalBuyingVolume - position.totalSellingVolume,
					buyQnt: position.buyQnt - position.sellingQnt
				}
			} else {
				pos = {
					totalSellingVolume: position.totalSellingVolume - position.totalBuyingVolume,
					sellingQnt: position.sellingQnt - position.buyQnt
				}
			}
			consolidateResult(results, st, pos, positionKey)
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

function consolidateResult(results, swingtrade, position, positionKey) {
	if (position.buyQnt > 0) { // compra
		if (swingtrade.quant >= 0 && swingtrade.total >= 0) {// comprado
			swingtrade.total += position.totalBuyingVolume
			swingtrade.quant += position.buyQnt
		} else {// vendido
			let avgPriceShort = swingtrade.total/swingtrade.quant
			let avgPriceLong = position.totalBuyingVolume/position.buyQnt
			if (position.buyQnt <= swingtrade.quant) {
				results.total += (avgPriceShort - avgPriceLong)*position.buyQnt
				swingtrade.quant += position.buyQnt
				swingtrade.total += position.totalBuyingVolume
			} else {
				console.log('tratar 1') // comprar posicao maior do que esta vendido
			}
		}
		results.volume += position.totalBuyingVolume
	} else if(position.sellingQnt > 0) { // venda
		if (swingtrade.quant >= 0 && swingtrade.total >= 0) {// comprado
			let avgPriceShort = position.totalSellingVolume/position.sellingQnt
			let avgPriceLong = swingtrade.total/swingtrade.quant
			if (position.sellingQnt <= swingtrade.quant) {
				results.total += (avgPriceShort - avgPriceLong)*position.sellingQnt
				swingtrade.quant -= position.sellingQnt
				swingtrade.total -= position.totalSellingVolume
			} else {
				results.total += (avgPriceShort - avgPriceLong)*swingtrade.quant
				const quantLeft = position.sellingQnt - swingtrade.quant
				swingtrade.quant = quantLeft
				swingtrade.total = quantLeft * avgPriceShort
			}
		} else {// vendido
			swingtrade.total -= position.totalSellingVolume
			swingtrade.quant += position.sellingQnt
		}
		results.volume += position.totalSellingVolume
		results.sold += position.totalSellingVolume
	}
}

function handleStockDivision(key, position) {
	if (key == 'TRAN PAULIST') {
		positionDate = new Date(position.date).getTime()
		stockDivisionDate = new Date('2019-04-04').getTime()
		if (positionDate <= stockDivisionDate) {
			position.buyQnt *= 4
			position.sellingQnt *= 4
		}
	}
}

function printResults(results) {
	Object.keys(results).forEach(k => {
		const r = results[k]
		console.log(`Período ${k}`)
		console.log(`Resultado Líquido: R$ ${(r.total - r.costs - r.irrf).toFixed(2).padStart(8, ' ')} ; Custos: R$ ${(r.costs).toFixed(2).padStart(8, ' ')} ; IRRF: R$ ${r.irrf.toFixed(2).padStart(8, ' ')}`)
	})
}