const form = document.querySelector('form')

form.addEventListener('submit', async e => {
    e.preventDefault()

    const formData = new FormData(e.target)

    try {
        await fetch('/api/product/store', {
            method: 'post',
            body: formData
        })
    } catch (e) {
        console.log(e)
        alert('Произошла ошибка при добавлении продукта')
    } finally {
        window.location.href = '/'
    }
})