function Login() {
  return (
    <div className="Login">
      <h1>Login</h1>
      <div>
        <label>
          Email
          <input type="text"></input>
        </label>
      </div>
      <div>
        <label>
          Password
          <input type="password"></input>
        </label>
      </div>
      <button>Login</button>
      <button>新規登録</button>
    </div>
  );
}

export default Login;
