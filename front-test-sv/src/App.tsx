import "@fortawesome/fontawesome-free/css/all.min.css";
import {
  BrowserRouter, Navigate, Route, Routes
} from "react-router-dom";
import './App.css';
import Add from "./pages/add";
import Edit from "./pages/edit";
import Post from "./pages/post";
import Preview from "./pages/preview";


function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={<Navigate to="/post" replace />}
        />
        <Route path="/post" element={<Post />} />
        <Route path="/post/add" element={<Add />} />
        <Route path="/post/preview" element={<Preview />} />
        <Route path="/post/edit/:id" element={<Edit />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
