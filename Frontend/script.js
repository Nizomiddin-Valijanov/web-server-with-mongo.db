const url = "http://localhost:8080/people";

async function postData(url = "", data = {}) {
  const response = await fetch(url, {
    method: "POST",
    credentials: "same-origin",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  return response.json();
}

let array = [
  {
    id: 1,
    name: "Nizomiddin",
    age: 16,
  },
  {
    id: 2,
    name: "SALOM OLAM",
    age: 22,
  },
  {
    id: 3,
    name: "Sherbek",
    age: 33,
  },
  {
    id: 4,
    name: "Jimi",
    age: 11,
  },
  {
    id: 5,
    name: "nizomi",
    age: 12,
  },
];

async function addPeople() {
  for (const item of array) {
    await postData(url, item);
  }
}

async function fetchData() {
  const data = await fetch(url);
  return data.json();
}

async function main() {
  await addPeople();
  const people = await fetchData();
  console.log(people);
}

main();
