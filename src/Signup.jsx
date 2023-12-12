import Login from "./Login";
import { Link } from "react-router-dom";

function Signup() {
  return (
    <div class="wrapper">
      <div class="container">
        <h1>Welcome</h1>

        <form className="form">
          <input type="text" placeholder="Name" />
          <input type="text" placeholder="Email" />
          <input type="password" placeholder="Password" />
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
