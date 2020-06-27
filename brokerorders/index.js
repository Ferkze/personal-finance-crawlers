const fs = require('fs');

let rawdata = fs.readFileSync('orders.json');
let orders = JSON.parse(rawdata);

let posicoesKeys = [...new Set(orders.map(o => o.date.substring(0, 10)))]

const posicoes = {}

orders.sort((a,b) => new Date(b.date) - new Date(a.date))
orders = orders.map(o => {
	o.date = o.date.substring(0, 10)
	return o
})

const futures = orders.filter(o => o.asset == 'WIN' || o.asset == 'WDO' || o.asset == 'DOL' || o.asset == 'IND')
posicoesKeys = [...new Set(futures.map(o => o.date.substring(0, 10)))]
console.log(`starting posicoesKeys.map`)
const resultados = posicoesKeys.map(k => {
	const trades = futures.filter(o => o.date == k)
	const daytrades = {
		date: k,
		WDO: {
			numberOfContracts: 0,
			result: 0,
			cost: 0
		},
		WIN: {
			numberOfContracts: 0,
			result: 0,
			cost: 0
		},
		DOL: {
			numberOfContracts: 0,
			result: 0,
			cost: 0
		},
		irrf: 0
	}
	for (const trade of trades) {
		if (trade.asset == 'WDO') {
			daytrades[trade.asset].result -= trade.price * trade.qnt * 10
		} else if (trade.asset == 'WIN') {
			daytrades[trade.asset].result -= trade.price * trade.qnt / 5
		} else if (trade.asset == 'DOL') {
			daytrades[trade.asset].result -= trade.price * trade.qnt * 50
		}
		daytrades[trade.asset].numberOfContracts += Math.abs(trade.qnt)
	}
	daytrades.WIN.cost = daytrades.WIN.numberOfContracts * 0.25
	daytrades.WDO.cost = daytrades.WDO.numberOfContracts * 0.91
	daytrades.DOL.cost = daytrades.DOL.numberOfContracts * 5 * 0.91

	const result = daytrades.WIN.result + daytrades.WDO.result + daytrades.DOL.result
	const cost = daytrades.WIN.cost + daytrades.WDO.cost + daytrades.DOL.cost

	if (result > 0 && result - cost > 0) {
		daytrades.irrf = (result - cost) * 0.01
	}

	return daytrades
})

console.log(`starting resultadosReduced`)
const resultadosReduced = resultados.reduce((acc, cur) => {
	const key = cur.date.substring(0,7)
	if (!acc[key]) {
		acc[key] = {
			WDO: cur.WDO.result - cur.WDO.cost + cur.DOL.result - cur.DOL.cost,
			WIN: cur.WIN.result - cur.WIN.cost,
			IRRF: cur.irrf
		}
	} else {
		acc[key].WDO += cur.WDO.result - cur.WDO.cost + cur.DOL.result - cur.DOL.cost
		acc[key].WIN += cur.WIN.result - cur.WIN.cost
		acc[key].IRRF += cur.irrf
	}
	return acc
}, {})

console.log(resultadosReduced);
