const searchParams = new URLSearchParams(window.location.search)
const productId = searchParams.get('id')

const name = document.querySelector('#name')
const price = document.querySelector('#price')
const composition = document.querySelector('#composition')
const photoUrl = document.querySelector('#photo_url')
const spinner = document.querySelector('.spinner-border')
const form = document.querySelector('form')

const renderForm = async () => {
    try {
        const res = await fetch(`/api/product/show?id=${productId}`)

        const product = await res.json()

        name.value = product.Name
        price.value = product.Price
        composition.value = product.Composition.join(', ')
        photoUrl.value = product.PhotoUrl
    } catch (e) {
        console.log(e)
        alert('Ошибка при получении продукта')
    } finally {
        spinner.classList.add('d-none')
        form.classList.remove('d-none')
    }
}

renderForm()

form.addEventListener('submit', async e => {
    e.preventDefault()

    const formData = new FormData(e.target)
    formData.append('id', productId)

    try {
        await fetch('/api/product/update', {
            method: 'post',
            body: formData
        })

        window.location.href = '/'
    } catch (e) {
        console.log(e)
        alert('Произошла ошибка при редактировании продукта')
    }
})