module.exports = {
	filterOrdersByType: (orders, type) => orders.filter(o => o.type == type),
	filterOrdersByDate: (orders, date) => orders.filter(o => o.date == date),
	sortOrdersByDate: (orders) => orders.sort((order1, order2) => new Date(order2.date) - new Date(order1.date))
}