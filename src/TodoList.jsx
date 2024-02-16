import SideMenu from "./components/SideMenu";
import add_list from "./image/add_list.png";
import download from "./image/download.png";
import rubbish from "./image/rubbish.png";
import React, { useState } from "react";

function TodoList() {
  const [inputData, setInputData] = useState([]);
  const [formData, setFormData] = useState({
    input: "",
    date: "",
  });

  const onSubmitForm = async () => {
    console.log(formData.input.value);
  };






  return (
    <div className="flex">
      <SideMenu />
      <div className="flex flex-col">
        <div className="flex mt-8 w-full ml-32 border-l-4 border-gray-300 pl-4">
          <h1 className="text-6xl font-bold mt-10 ml-4 mr-72">Todoリスト</h1>
          <span className="ml-auto">
            <button className="bg-yellow-300 rounded-lg">
              <img className="w-20" src={add_list} />
            </button>
            <button className="ml-11 bg-green-300 rounded-lg">
              <img className="w-20" src={download} />
            </button>
            <button className="ml-10 mr-10 bg-red-400 rounded-lg">
              <img className="w-20" src={rubbish} />
            </button>
          </span>
        </div>
        <div className="ml-32 mt-4 gird">
          <form>
            <label className="text-2xl">
              やること
              <input
                className="border-3 border-slate-600 ml-0 w-full text-left"
                type="text"
                value={formData.input}
                onChange={(e) =>
                  setFormData({ ...formData, input: e.target.value })
                }
              />
            </label>
            <label className="text-2xl">
              期日
              <input
                className="border-3 border-slate-600 ml-0 w-full"
                type="date"
                value={formData.date}
                onChange={(e) =>
                  setFormData({ ...formData, date: e.target.value })
                }
              />
            </label>
            <input
              className="text-slate-950 bg-emerald-300 hover:bg-emerald-500 rounded-lg mt-10"
              onClick={onSubmitForm}
              type="submit"
              value="登録"
            />
          </form>
          <div>
            <input className="text-xl" type="checkbox"/>
            
          </div>
        </div>
      </div>
    </div>
  );
}

export default TodoList;
