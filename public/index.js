console.log("start");

async function sendRequest(i) {
  const res = await fetch(`http://localhost:9999/api/?n=${i}`);
  console.log(res.text());
}

const arr = [...Array(10)].map((v, i) => i);

Promise.all(arr.map(sendRequest));
