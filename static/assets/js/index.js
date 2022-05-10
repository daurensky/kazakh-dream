const spinner = document.querySelector('.spinner-border')
const productsTableWrapper = document.querySelector('#products-table-wrapper')
const productsTableBody = productsTableWrapper.querySelector('tbody')

const renderOrders = async () => {
    try {
        const res = await fetch('/api/product')

        const products = await res.json()

        products?.forEach(({Id, Name, Price, PhotoUrl, Composition}, i) => {
            productsTableBody.innerHTML += `
                <tr>
                    <td>${i + 1}</td>
                    <td>${Name}</td>
                    <td>${Price} тг.</td>
                    <td>
                        <img src="${PhotoUrl}" alt="Product" width="100" class="rounded">
                    </td>
                    <td>${Composition.join(', ')}</td>
                    <td>
                        <div class="d-flex justify-content-end gap-1">
                            <a href="product/edit.html?id=${Id}" class="btn btn-outline-success">Редактировать</a>
                        </div>
                    </td>
                </tr>
            `
        })
    } catch (e) {
        console.log(e)
        alert('Произошла ошибка при загрузки меню')
    } finally {
        spinner.classList.add('d-none')
        productsTableWrapper.classList.remove('d-none')
    }
}

renderOrders()