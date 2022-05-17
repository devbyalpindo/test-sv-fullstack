/*eslint-disable*/
import React from "react";
import { Link, useLocation } from "react-router-dom";

interface NavlinkProps {
  name: string;
  url: string;
  icon: string
}

const Navlink: React.FC<NavlinkProps> = ({ name, url, icon }) => {
  const location = useLocation();
  const { pathname } = location;

  return <li className="items-center">
    <Link
      className={
        "text-xs uppercase py-3 font-bold block " +
        (pathname === url
          ? "text-blue-500 hover:text-blue-600"
          : "text-gray-700 hover:text-gray-500")
      }
      to={url}
    >
      <i
        className={
          `fas ${icon} mr-2 text-sm ` +
          (pathname === url
            ? "opacity-75"
            : "text-gray-300")
        }
      ></i>{" "}
      {name}
    </Link>
  </li>
}


const Sidebar = () => {
  const [collapseShow, setCollapseShow] = React.useState("hidden");
  return (
    <>
      <nav className="md:left-0 md:block md:fixed md:top-0 md:bottom-0 md:overflow-y-auto md:flex-row md:flex-nowrap md:overflow-hidden shadow-xl bg-white flex flex-wrap items-center justify-between relative md:w-64 z-10 py-4 px-6">
        <div className="md:flex-col md:items-stretch md:min-h-full md:flex-nowrap px-0 flex flex-wrap items-center justify-between w-full mx-auto">
          {/* Toggler */}
          <button
            className="cursor-pointer text-black opacity-50 md:hidden px-3 py-1 text-xl leading-none bg-transparent rounded border border-solid border-transparent"
            type="button"
            onClick={() => setCollapseShow("bg-white m-2 py-3 px-6")}
          >
            <i className="fas fa-bars"></i>
          </button>
          {/* Brand */}
          <Link
            className="md:block text-left md:pb-2 text-gray-600 mr-0 inline-block whitespace-nowrap text-sm uppercase font-bold p-4 px-0"
            to="/"
          >
            Article Test SV
          </Link>
          {/* Collapse */}
          <div
            className={
              "md:flex md:flex-col md:items-stretch md:opacity-100 md:relative md:mt-4 md:shadow-none shadow absolute top-0 left-0 right-0 z-40 overflow-y-auto overflow-x-hidden h-auto items-center flex-1 rounded " +
              collapseShow
            }
          >
            {/* Collapse header */}
            <div className="md:min-w-full md:hidden block pb-4 mb-4 border-b border-solid border-gray-200">
              <div className="flex flex-wrap">
                <div className="w-6/12">
                  <Link
                    className="md:block text-left md:pb-2 text-gray-600 mr-0 inline-block whitespace-nowrap text-sm uppercase font-bold p-4 px-0"
                    to="/"
                  >
                    Article Test SV
                  </Link>
                </div>
                <div className="w-6/12 flex justify-end">
                  <button
                    type="button"
                    className="cursor-pointer text-black opacity-50 md:hidden px-3 py-1 text-xl leading-none bg-transparent rounded border border-solid border-transparent"
                    onClick={() => setCollapseShow("hidden")}
                  >
                    <i className="fas fa-times"></i>
                  </button>
                </div>
              </div>
            </div>


            {/* Divider */}
            <hr className="my-4 md:min-w-full" />
            {/* Heading */}
            <h6 className="md:min-w-full text-gray-500 text-xs uppercase font-bold block pt-1 pb-4 no-underline">
              Post
            </h6>
            {/* Navigation */}

            <ul className="md:flex-col md:min-w-full flex flex-col list-none">
              <Navlink name="All Post" icon="fa-newspaper" url="/post" />
              <Navlink name="Add New" icon="fa-square-plus" url="/post/add" />
              <Navlink name="Preview" icon="fa-eye" url="/post/preview" />
            </ul>
          </div>
        </div>
      </nav>
    </>
  );
}

export default Sidebar;