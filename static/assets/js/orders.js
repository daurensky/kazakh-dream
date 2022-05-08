const spinner = document.querySelector('.spinner-border')
const ordersTable = document.querySelector('#orders-table')
const ordersTableBody = ordersTable.querySelector('tbody')

const statuses = [{
    id: 'PREPARING', title: 'Готовится'
}, {
    id: 'SENT', title: 'Отправлено курьером'
}, {
    id: 'DELIVERED', title: 'Доставлено'
}]

const renderOrders = async () => {
    spinner.classList.remove('d-none')
    ordersTable.classList.add('d-none')
    ordersTableBody.innerHTML = ''

    try {
        const res = await fetch('/order')

        const orders = await res.json()

        orders.forEach(({Id, Client, CreatedAt, Products, Status}, i) => {
            let totalPrice = 0
            let composition = []

            Products.forEach(({Price, Name}) => {
                totalPrice += Price
                composition.push(Name)
            })

            ordersTableBody.innerHTML += `
                <tr>
                    <td>${i + 1}</td>
                    <td>${Client.Address}, ${Client.Phone}, ${Client.Name}</td>
                    <td>${CreatedAt}</td>
                    <td>${totalPrice} тг.</td>
                    <td>${composition.join(', ')}</td>
                    <td>
                        <div class="d-flex gap-1">
                            <select class="form-select w-auto order-status" data-order-id="${Id}">
                                ${statuses.map(({id, title}) => (`
                                    <option value="${id}" ${id === Status ? 'selected' : ''}>${title}</option>
                                `))}
                            </select>
                            <button class="btn btn-outline-primary flex-shrink-0 order-status-btn"
                                data-order-id="${Id}">Поменять статус</button>
                        </div>
                    </td>
                </tr>
            `
        })

        document.querySelectorAll('.order-status-btn').forEach(btn => {
            btn.addEventListener('click', () => {
                const orderId = btn.dataset.orderId
                const status = document.querySelector(`.order-status[data-order-id="${orderId}"]`).value
                updateOrderStatus(orderId, status)
            })
        })
    } catch (e) {
        alert('Произошла ошибка при загрузки меню')
    } finally {
        spinner.classList.add('d-none')
        ordersTable.classList.remove('d-none')
    }
}

renderOrders()

const updateOrderStatus = async (orderId, status) => {
    try {
        const formData = new FormData()
        formData.append('order_id', orderId)
        formData.append('status', status)

        await fetch('/update-order-status', {
            method: 'post',
            body: formData
        })

        alert('Статус успешно изменен!')
    } catch (e) {
        alert('Произошла ошибка при изменентт статуса')
    }
}