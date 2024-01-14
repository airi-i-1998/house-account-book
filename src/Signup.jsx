import React, { useState } from "react"
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";

function Signup() {

  //配列に分割代入している
  //array[0],array[1]
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();
  const [errorMessage, setErrorMessage] = useState("");

  //フォームが送信されてもページが再読み込みされず、非同期処理が行われる
  const handleSignup = async (e) => {
    //フォームのデフォルトの送信動作がキャンセルされる→フォームが送信されるとページが再読み込みされるのを防ぐ
    e.preventDefault();

    const formData = {
      name: name,
      email: email,
      password: password,
    };

    try {
      // APIエンドポイントにデータを送信するためのコードを記述
      const response = await fetch("http://localhost:8080/signup", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      // レスポンスの処理
      if (response.ok) {
        navigate('/Home');
      } else {
        //APIからのレスポンスをjson形式に解析
        const data = await response.json();
        if (data && data.error) {
          setErrorMessage(data.error);
        } else {
          // エラーメッセージの表示
          throw new Error("Sign up failed");
        }
      }
    } catch (error) {
      console.error("Error:", error);
    }
  };

  return (
    <div class="wrapper">
      <div class="container">
        <h1>Welcome</h1>
        <div className="error-message">
        {errorMessage.split('\n').map((error, index) => (
            <div key={index}>{error}</div>
          ))}
          </div>
        <form className="form" onSubmit={handleSignup}>
          <input type="text" placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} />
          <input type="text" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
          <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
          <button type="submit" id="login-button">Sign up</button>
          <p>Already have an account?　<Link to="/Login">Login</Link></p>
        </form>
      </div>

      <ul className="bg-bubbles">
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
      </ul>
    </div>

  );
}

export default Signup;
