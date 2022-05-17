import { useEffect, useState } from "react";
import { useSearchParams } from "react-router-dom";
import swal from "sweetalert";
import { getAllArticleWithLimit } from "../api/article";
import { ArticleTypes } from "../api/data-type/article";
import Layout from "../layout";

const Preview = () => {
  let [searchParams] = useSearchParams();
  const [pages, setPages] = useState<number>(
    searchParams.get("page") ? parseInt(searchParams.get("page")!) : 1
  );
  const [limit, setLimit] = useState(3);
  const [total, setTotal] = useState(0);
  const [offset, setOffset] = useState(0);
  const [article, setArticle] = useState<ArticleTypes[]>([]);
  const [pagination, setPagination] = useState([1, 2, 3, 4, 5]);

  const getArticle = async (limit: number, offset: number) => {
    const response = await getAllArticleWithLimit(limit, offset);

    if (response.statusCode === 200) {
      setTotal(response.data.total);
      setArticle(response.data.article);
      generatePagination(response.data.total);
    } else {
      swal("Oops", response.error, "error");
    }
  };

  const generatePagination = (total: number) => {
    let paging = [];
    for (let i = 1; i <= Math.ceil(total / limit); i++) {
      paging.push(i);
    }
    setPagination(paging);
  };

  const changePage = (page: number) => {
    const offsetData = page * limit - limit;
    setOffset(offsetData);
    setPages(page);
  };

  useEffect(() => {
    getArticle(limit, offset);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [pages, limit]);

  return (
    <Layout>
      <div className='p-10'>
        <h1 className='text-2xl font-bold text-slate-700'>Preview</h1>
        <div className='flex items-center space-x-2 mt-10'>
          <p>Show :</p>
          <select
            onChange={(e) => {
              changePage(1);
              setLimit(parseInt(e.target.value));
            }}
            className='border p-1'>
            <option value={3}>3</option>
            <option value={5}>5</option>
            <option value={10}>10</option>
          </select>
          <p>of {total}</p>
        </div>

        <div className='mt-10 grid grid-cols-4 gap-4'>
          {article.map((item) => {
            return (
              <div
                key={item.id}
                className='p-2 border border-slate-500 rounded-md shadow-md'>
                <h2 className='font-bold text-lg text-slate-700'>
                  {item.title}
                </h2>
                <p className='mt-4 text-slate-500'>
                  category : {item.category}
                </p>
                <p className='mt-4 text-md text-slate-700 line-clamp-5'>
                  {item.content}
                </p>
              </div>
            );
          })}
        </div>
        <div className='w-full flex space-x-2 mt-5 justify-center'>
          <button
            onClick={() => changePage(pages - 1)}
            disabled={pages === 1}
            className={`${pages === 1 ? "text-slate-300" : "text-slate-700"}`}>
            previous
          </button>
          {pagination.map((index: number) => {
            return index === pages ? (
              <p className='text-blue-500'>{index}</p>
            ) : (
              <button
                className='text-slate-700'
                onClick={() => changePage(index)}>
                {index}
              </button>
            );
          })}
          <button
            onClick={() => changePage(pages + 1)}
            disabled={pages === pagination.at(-1)}
            className={`${
              pages === pagination.at(-1) ? "text-slate-300" : "text-slate-700"
            }`}>
            next
          </button>
        </div>
      </div>
    </Layout>
  );
};

export default Preview;
