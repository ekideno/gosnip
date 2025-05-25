document.addEventListener('DOMContentLoaded', () => {
    const addSnippetForm = document.getElementById('addSnippetForm');
    const resultSection = document.getElementById('resultSection');

    addSnippetForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        
        const contentInput = document.getElementById('content');
        const errorMessage = contentInput.parentElement.nextElementSibling;
        
        contentInput.classList.remove('error');
        errorMessage.style.display = 'none';
        
        const content = contentInput.value.trim();
        const customCode = document.getElementById('customCode').value.trim();
        
        if (!content) {
            contentInput.classList.add('error');
            errorMessage.style.display = 'flex';
            return;
        }
        
        try {
            const response = await fetch('http://localhost:9000/random', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    original_url: content
                })
            });
            
            if (!response.ok) {
                const error = await response.json();
                console.error('Server error:', error);
                contentInput.classList.add('error');
                errorMessage.querySelector('span').textContent = error.error || 'An error occurred';
                errorMessage.style.display = 'flex';
                return;
            }
            
            const data = await response.json();
            displayResult({
                original_url: data.url,
                slug: data.slug
            });
        } catch (error) {
            contentInput.classList.add('error');
            errorMessage.querySelector('span').textContent = 'Failed to connect to server';
            errorMessage.style.display = 'flex';
        }
    });
});

function displayResult(data) {
    const resultSection = document.getElementById('resultSection');
    const originalUrl = document.getElementById('originalUrl');
    const shortenedUrl = document.getElementById('shortenedUrl');
    const urlCode = document.getElementById('urlCode');

    document.querySelectorAll('.copied').forEach(el => el.classList.remove('copied'));
    document.querySelectorAll('.copy-icon').forEach(icon => icon.classList.remove('show'));

    originalUrl.textContent = data.original_url;
    shortenedUrl.textContent = `http://gosnip/${data.slug}`;
    urlCode.textContent = data.slug;

    resultSection.style.display = 'block';
    resultSection.scrollIntoView({ behavior: 'smooth' });
}

async function copyToClipboard(element) {
    try {
        const text = element.textContent.trim();
        
        const textArea = document.createElement('textarea');
        textArea.value = text;
        textArea.style.position = 'fixed';
        textArea.style.left = '-999999px';
        textArea.style.top = '-999999px';
        document.body.appendChild(textArea);
        textArea.focus();
        textArea.select();
        
        try {
            document.execCommand('copy');
            
            document.querySelectorAll('.copied').forEach(el => el.classList.remove('copied'));
            document.querySelectorAll('.copy-icon').forEach(icon => icon.classList.remove('show'));
            
            element.classList.add('copied');
            const icon = element.querySelector('.copy-icon');
            if (icon) {
                icon.classList.add('show');
            }
            
            setTimeout(() => {
                element.classList.remove('copied');
                if (icon) {
                    icon.classList.remove('show');
                }
            }, 2000);
        } catch (err) {
            console.error('Fallback: Oops, unable to copy', err);
        }
        
        textArea.remove();
    } catch (err) {
        console.error('Failed to copy text: ', err);
    }
}

async function deleteSnippet(id) {
    if (!confirm('Are you sure you want to delete this URL?')) {
        return;
    }

    try {
        const response = await fetch(`/api/snippets/${id}`, {
            method: 'DELETE'
        });

        if (response.ok) {
            loadSnippets();
        } else {
            alert('Error deleting URL');
        }
    } catch (error) {
        console.error('Error deleting URL:', error);
        alert('An error occurred while deleting the URL');
    }
} 