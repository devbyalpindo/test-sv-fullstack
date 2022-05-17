import { Link } from "react-router-dom";
import swal from "sweetalert";
import { removeArticleById } from "../../api/article";
import { ArticleTypes } from "../../api/data-type/article";

interface TrashedProps {
    data: ArticleTypes[];
    update: boolean;
    setUpdate: React.Dispatch<React.SetStateAction<boolean>>
}

const Trashed: React.FC<TrashedProps> = ({ data, update, setUpdate }) => {

    const removeArticle = async (id: number) => {
        swal({
            title: "Are you sure?",
            text: "Once delete article, you will not be able to revert this change!",
            icon: "warning",
            buttons: ["cancel", "Delete"],
            dangerMode: true,
        }).then(async (willdelete) => {
            if (willdelete) {
                const response = await removeArticleById(id);

                if (response.statusCode === 200) {
                    swal("success", "article deleted", "success")
                    setUpdate(!update);
                } else {
                    swal("Oops", response.error, "error");
                }
            }
        });

    }

    return <table className="mt-10 border-collapse table-auto w-full text-sm">
        <thead className="bg-white">
            <tr>
                <th className="border-b font-medium p-4 pl-8 pt-0 pb-3 text-slate-700 text-left">Title</th>
                <th className="border-b font-medium p-4 pt-0 pb-3 text-slate-700 text-left">Category</th>
                <th className="border-b font-medium p-4 pr-8 pt-0 pb-3 text-slate-700 text-left">Action</th>
            </tr>
        </thead>
        <tbody className="bg-white">
            {data.map((article) => {
                return <tr key={article.id}>
                    <td className="border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-slate-500 dark:text-slate-400">{article.title}</td>
                    <td className="border-b border-slate-100 dark:border-slate-700 p-4 text-slate-500 dark:text-slate-400">{article.category}</td>
                    <td className="border-b border-slate-100 dark:border-slate-700 p-4 pr-8 text-slate-500 dark:text-slate-400">
                        <div className="flex space-x-5">
                            <Link to={`/post/edit/${article.id}`}><i className="fas fa-pencil text-blue-500"></i></Link>
                            <button onClick={() => removeArticle(article.id)}><i className="fas fa-trash text-red-500"></i></button>
                        </div>
                    </td>
                </tr>
            })}

        </tbody>
    </table>
}

export default Trashed;