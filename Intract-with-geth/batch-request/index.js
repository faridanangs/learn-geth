
async function main() {
    const endpoint = 'http://127.0.0.1:3334';
  
    const res = await fetch(endpoint, {
      method: 'POST',
      body: JSON.stringify({
          id: 0x1,
        method: 'eth_getBalance',
        params: ["0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326", "latest"],
        jsonrpc: '2.0'
      }),
      headers: { 'Content-Type': 'application/json' }
    });
    const data = await res.json();
    console.log(data)
  }
  
  main()
    .then()
    .catch(err => console.log(err));