import React from "react";
import icon from "../image/icon.png";
import home from "../image/home.png";
import money from "../image/money.png";
import calendar from "../image/calendar.png";
import memo from "../image/memo.png";
import { useNavigate } from 'react-router-dom';

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
      <ul>
        <span className="list">
          <img src={home} className="home-icon" />
          ホーム
        </span>
        <span className="list">
          <img src={money} className="home-icon" />
          入出金
        </span>
        <span className="list">
          <img src={calendar} className="home-icon" />
          月別の管理
        </span>
        <span className="list">
          <img src={memo} className="home-icon" />
          メモリスト
        </span>
      </ul>
      <div className="logout-wapper">
        <button className="logout" onClick={handleLogout}>ログアウト</button>
      </div>
    </div>
  );
}

export default SideMenu;