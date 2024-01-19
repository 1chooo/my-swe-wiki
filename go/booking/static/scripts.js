// 透過 JavaScript 與後端進行通訊
async function getBooks() {
    const response = await fetch('/books');
    const books = await response.json();

    const bookList = document.getElementById('bookList');
    bookList.innerHTML = '';

    books.forEach(book => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${book.ID}</td>
            <td>${book.Title}</td>
            <td>${book.Author}</td>
            <td><button onclick="deleteBook('${book.ID}')">刪除</button></td> <!-- 新增刪除按鈕 -->
        `;
        bookList.appendChild(row);
    });
}

async function addBook(event) {
    event.preventDefault();

    const id = document.getElementById('id').value;
    const title = document.getElementById('title').value;
    const author = document.getElementById('author').value;

    const response = await fetch('/books', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            ID: id,
            Title: title,
            Author: author,
        }),
    });

    if (response.ok) {
        getBooks();
    } else {
        console.error('新增書本失敗');
    }
}

async function deleteBook(bookId) {
    const response = await fetch(`/books/${bookId}`, {
        method: 'DELETE',
    });

    if (response.ok) {
        getBooks();
    } else {
        console.error('刪除書本失敗');
    }
}

// 初次載入頁面時獲取書本列表
document.addEventListener('DOMContentLoaded', () => {
    getBooks();
});

// 監聽新增書本表單的提交事件
const addBookForm = document.getElementById('addBookForm');
addBookForm.addEventListener('submit', addBook);
