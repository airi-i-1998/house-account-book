import React from "react";
import SideMenu from './components/SideMenu';
import add_list from "./image/add_list.png";
import download from "./image/download.png";
import rubbish from "./image/rubbish.png";

function TodoList(){
  return(
    <div className="flex">
      <SideMenu />
      <div className="flex mt-8 w-full ml-32">
      <h1 className="text-6xl font-bold mt-10 ml-4">やることリスト</h1>
      <span className="ml-auto">
        <button className="bg-yellow-300 rounded-lg">
          <img className="w-20" src={add_list} />
        </button>
        <button  className="ml-11 bg-green-300 rounded-lg">
          <img className="w-20" src={download} />
        </button>
        <button className="ml-10 mr-10 bg-red-400 rounded-lg">
          <img className="w-20" src={rubbish} />
        </button>
      </span>
      </div>
    </div>
  );
}

export default TodoList;
