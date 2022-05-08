const form = document.querySelector('form')

form.addEventListener('submit', async e => {
    e.preventDefault()

    const formData = new FormData(e.target)

    try {
        await fetch('/product/store', {
            method: 'post',
            body: formData
        })

        window.location.href = '/'
    } catch (e) {
        alert('Произошла ошибка при добавлении продукта')
    }
})