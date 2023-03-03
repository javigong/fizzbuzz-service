import { useEffect, useState } from "react";
import "./App.css";

function App() {
  const fizzbuzzUrl = import.meta.env.VITE_FIZZBUZZ_URL as string;
  const [count, setCount] = useState(0);
  const [message, setMessage] = useState("");
  console.log({ fizzbuzzUrl });
  const postCountData = {
    count: count,
  };

  useEffect(() => {
    fetch(`${fizzbuzzUrl}/fizzbuzz`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(postCountData),
    })
      .then((response) => response.json())
      .then((data) => {
        setMessage(data.message);
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }, [count]);

  return (
    <div className="App">
      <div className="flex flex-col min-h-screen overflow-auto justify-center gap-16">
        <section>
          <h1 className="text-lg font-light tracking-widest">Your count</h1>
          <p className="text-lg font-light">{count}</p>
        </section>
        <div>
          <button
            className="text-white text-lg font-light tracking-wide bg-blue-600 px-9 py-3 rounded-md shadow-md shadow-slate-400 transition duration-150 active:scale-95"
            onClick={() => setCount((count) => count + 1)}
          >
            Push me!
          </button>
        </div>
        <div className="h-[48px]">
          <p className="text-5xl font-medium font-roboto-mono text">{message}</p>
        </div>
      </div>
    </div>
  );
}

export default App;
