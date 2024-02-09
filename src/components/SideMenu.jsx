import React from "react";
import icon from "../image/icon.png";
import home from "../image/home.png";
import money from "../image/money.png";
import calendar from "../image/calendar.png";
import memo from "../image/memo.png";
import { useNavigate } from 'react-router-dom';
import { Link } from "react-router-dom";

const SideMenu = () => {
  const navigate = useNavigate();

  const handleLogout = () => {
    navigate('/');
  };

  return (
    <div className="side-menu">
      <h1 className="icon-wrapper">
        <img src={icon} className="icon" />
      </h1>
      <ul className="ml-3.5">
        <span className="list">
          <img src={home} className="home-icon" />
          <Link to="/Home">ホーム</Link>
        </span>
        <span className="list">
          <img src={money} className="home-icon" />
          <Link to="/Balance">入出金</Link>
        </span>
        <span className="list">
          <img src={calendar} className="home-icon" />
          <Link to="/Calendar">月別の管理</Link>
        </span>
        <span className="list">
          <img src={memo} className="home-icon" />
          <Link to="/TodoList">メモリスト</Link>
        </span>
      </ul>
      <div className="logout-wapper">
        <button className="logout" onClick={handleLogout}>ログアウト</button>
      </div>
    </div>
  );
}

export default SideMenu;
