export const api = {
  async post(url, data) {
    const res = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': localStorage.getItem('token') ? `Bearer ${localStorage.getItem('token')}` : undefined
      },
      body: JSON.stringify(data)
    });
    return res.json();
  }
};
