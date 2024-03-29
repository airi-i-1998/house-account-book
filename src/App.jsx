import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Login from "./Login";
import Signup from "./Signup";
import Calendar from "./Calendar";
import Balance from "./Balance";
import TodoList from "./TodoList";

function App() {
  return (
    <BrowserRouter basename={process.env.PUBLIC_URL}>
      <Routes>
        <Route path="/" element={<Signup />} />
        <Route path="/Login" element={<Login />} />
        <Route path="TodoList" element={<TodoList />} />
        <Route path="/Calendar" element={<Calendar />} />
        <Route path="/Balance" element={<Balance />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
