import React, { useState } from "react";
import SideMenu from './components/SideMenu';

function Balance() {
  const [inputData, setInputData] = useState([]);
  const [formData, setFormData] = useState({
    description: '',
    date: '2024-02-01',
    amount: '',
    category: '食費',
    memo: '',
  });

  // 2. データの表示
  const inputList = inputData.map((data, index) => (
    <div key={index} className="mt-4 text-2xl w-full flex">
      <span className="w-2/12">{data.date}</span>
      <span className="w-2/12">{data.description}</span>
      <span className="w-2/12">{data.amount}</span>
      <span className="w-2/12">{data.category}</span>
      <span className="w-3/12">{data.memo}</span>
      <button onClick={() => handleDash(index)} className="w-1/12 mr-4">削除</button>
    </div>
  ));
  console.log(inputList);


  // 3. 高さの指定とスクロール
  const inputListContainerStyle = {
    maxHeight: '300px',  // 適切な高さを指定
    overflowY: 'auto',
  };

  const handleChange = async (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: name === 'amount' ? Number(value) : value,
    }));
  };

  const handleRegistration = async () => {
    // 登録ボタンがクリックされたときの処理
    // データをinputDataに追加するなどの処理を実装
    setInputData((prevData) => [
      ...prevData,
      {
        date: formData.date,
        description: formData.description,
        amount: formData.amount,
        category: formData.category,
        memo: formData.memo,
      },
    ]);

    try {
      const response = await fetch("http://localhost:8080/balance", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          description: formData.description,
          date: formData.date,
          amount: formData.amount,
          category: formData.category,
          memo: formData.memo,
        }),
      });
      if (response.ok) {
        const data = await response.json();
        // データを処理するなど
        console.log("Data from server:", data);
      } else {
        console.error("Error:", response.statusText);
      }
    } catch (error) {
      console.error("Error:", error);
    }

  };

  const handleDash = (index) => {
    setInputData((prevData) => {
      const newData = [...prevData];
      newData.splice(index, 1);
      return newData;
    });
  };

  return (
    <div className="flex h-full mt-8">
      <SideMenu />
      <div className="w-full ml-40 mt-8 ">
        <h1 className="text-5xl underline underline-offset-4">入力</h1>
        <div className='mt-8 text-2xl'>
          <input
            type="radio"
            id="income"
            name="description"
            value="収入"  // 収入の場合の値
            checked={formData.description === '収入'}
            onChange={handleChange}
            className='mr-3 items-center'
          />
          <label htmlFor="income" className='mr-4'>収入</label>
          <input
            type="radio"
            id="expense"
            name="description"
            value="支出"  // 支出の場合の値
            checked={formData.description === '支出'}
            onChange={handleChange}
            className="mr-3 items-center"
          />
          <label htmlFor="expense">支出</label>
          <label className="mt-6 flex items-center">
            日付
            <input
              type="date"
              id="date"
              name="date"
              value={formData.date}
              onChange={handleChange}
              className="ml-3"
            />
          </label>
          <span className="mt-6 flex items-center">
            金額
            <input
              type="number"
              id="amount"
              name="amount"
              value={formData.amount}
              onChange={handleChange}
              className="ml-3 text-lg text-right"
            />円
          </span>
          <span className="mt-6 flex items-center">
            カテゴリ
            <select name="category"
              value={formData.category}
              onChange={handleChange}
              className="ml-3 text-lg"
            >
              <option>食費</option>
              <option>日用品費</option>
              <option>医療費</option>
              <option>美容費</option>
              <option>交際費</option>
              <option>交通費</option>
              <option>娯楽費</option>
              <option>雑費</option>
              <option>特別費</option>
              <option>その他</option>
            </select>
          </span>
          <span className="mt-6 flex items-center">
            メモ
            <textarea
              name="memo"
              value={formData.memo}
              onChange={handleChange}
              className="ml-3 text-lg"
            />
          </span>
        </div>
        <button onClick={handleRegistration} className="mt-6 text-lg p-1 bg-emerald-300 hover:bg-emerald-700 rounded-full">登録する</button>
        <div className="w-full" style={inputListContainerStyle}>
          <h1 className="mt-10 text-5xl underline underline-offset-4">入力一覧</h1>
          {inputList}
        </div>
      </div>
    </div>

  );
}

export default Balance;
