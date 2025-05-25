// static/js/main.js
document.addEventListener('DOMContentLoaded', () => {
    const addSnippetForm = document.getElementById('addSnippetForm');
    const resultSection = document.getElementById('resultSection');
  
    addSnippetForm.addEventListener('submit', async (e) => {
      e.preventDefault();
  
      const contentInput = document.getElementById('content');
      const contentError = contentInput.parentElement.nextElementSibling;
  
      const customInput = document.getElementById('customCode');
      const customError = customInput.nextElementSibling;
  
      // Убираем старые ошибки
      contentInput.classList.remove('error');
      contentError.style.display = 'none';
      customInput.classList.remove('error');
      customError.style.display = 'none';
  
      const content = contentInput.value.trim();
      const customSlug = customInput.value.trim();
  
      let hasError = false;
  
      // Валидация URL
      if (!content) {
        contentInput.classList.add('error');
        contentError.querySelector('span').textContent = 'Please enter a valid URL';
        contentError.style.display = 'flex';
        hasError = true;
      }
  
      // Валидация custom code: только a–z
      if (customSlug && !/^[a-z]+$/.test(customSlug)) {
        customInput.classList.add('error');
        customError.querySelector('span').textContent = 'Custom code must contain only lowercase English letters (a–z)';
        customError.style.display = 'flex';
        hasError = true;
      }
  
      if (hasError) return;
  
      // Отправляем на сервер
      try {
        const requestBody = { original_url: content };
        if (customSlug) requestBody.slug = customSlug;
  
        const response = await fetch('http://localhost:9000/snip', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(requestBody)
        });
  
        if (!response.ok) {
          const error = await response.json();
          contentInput.classList.add('error');
          contentError.querySelector('span').textContent = error.error || 'An error occurred';
          contentError.style.display = 'flex';
          return;
        }
  
        const data = await response.json();
        displayResult({
          original_url: data.url,
          slug: data.slug
        });
      } catch (err) {
        contentInput.classList.add('error');
        contentError.querySelector('span').textContent = 'Failed to connect to server';
        contentError.style.display = 'flex';
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
    const text = element.textContent.trim();
    const textArea = document.createElement('textarea');
    textArea.value = text;
    textArea.style.position = 'fixed';
    textArea.style.left = '-9999px';
    document.body.appendChild(textArea);
    textArea.focus();
    textArea.select();
    try {
      document.execCommand('copy');
      document.querySelectorAll('.copied').forEach(el => el.classList.remove('copied'));
      document.querySelectorAll('.copy-icon').forEach(icon => icon.classList.remove('show'));
      element.classList.add('copied');
      const icon = element.querySelector('.copy-icon');
      if (icon) icon.classList.add('show');
      setTimeout(() => {
        element.classList.remove('copied');
        if (icon) icon.classList.remove('show');
      }, 2000);
    } finally {
      document.body.removeChild(textArea);
    }
  }
  
  async function deleteSnippet(id) {
    if (!confirm('Are you sure you want to delete this URL?')) return;
    try {
      const response = await fetch(`/api/snippets/${id}`, { method: 'DELETE' });
      if (response.ok) loadSnippets();
      else alert('Error deleting URL');
    } catch (err) {
      alert('An error occurred while deleting the URL');
    }
  }
  