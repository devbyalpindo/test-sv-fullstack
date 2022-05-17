import { Tab } from "@headlessui/react";
import { useEffect, useState } from "react";
import { getAllArticle } from "../api/article";
import { ArticleTypes } from "../api/data-type/article";
import Draft from "../components/Post/Draft";
import Published from "../components/Post/Published";
import Trashed from "../components/Post/Trashed";
import Layout from "../layout";

function classNames(...classes: any) {
    return classes.filter(Boolean).join(' ')
}

const Post = () => {
    const [published, setPublish] = useState<ArticleTypes[]>([]);
    const [draft, setDraft] = useState<ArticleTypes[]>([]);
    const [trashed, setTrashed] = useState<ArticleTypes[]>([]);
    const [update, setUpdate] = useState(false);

    const getArticle = async () => {
        const response = await getAllArticle();
        if (response.statusCode === 200) {
            const publishData = response.data.filter((article: ArticleTypes) => article.status === "Published");
            const draftData = response.data.filter((article: ArticleTypes) => article.status === "Draft");
            const trashedData = response.data.filter((article: ArticleTypes) => article.status === "Trashed");
            setDraft(draftData)
            setPublish(publishData)
            setTrashed(trashedData)
        }
    }

    useEffect(() => {
        getArticle();
    }, [update])

    return <Layout><div className="w-full px-2 py-16">
        <Tab.Group>
            <Tab.List className="flex space-x-1 rounded-xl bg-blue-900/20 p-1">
                <Tab
                    className={({ selected }) =>
                        classNames(
                            'w-full rounded-lg py-2.5 text-sm font-medium leading-5 text-blue-700',
                            'ring-white ring-opacity-60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2',
                            selected
                                ? 'bg-white shadow'
                                : 'text-blue-500 hover:bg-white/[0.12] hover:text-white'
                        )
                    }
                >
                    Published ({published.length})
                </Tab>
                <Tab
                    className={({ selected }) =>
                        classNames(
                            'w-full rounded-lg py-2.5 text-sm font-medium leading-5 text-blue-700',
                            'ring-white ring-opacity-60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2',
                            selected
                                ? 'bg-white shadow'
                                : 'text-blue-500 hover:bg-white/[0.12] hover:text-white'
                        )
                    }
                >
                    Draft ({draft.length})
                </Tab>
                <Tab
                    className={({ selected }) =>
                        classNames(
                            'w-full rounded-lg py-2.5 text-sm font-medium leading-5 text-blue-700',
                            'ring-white ring-opacity-60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2',
                            selected
                                ? 'bg-white shadow'
                                : 'text-blue-500 hover:bg-white/[0.12] hover:text-white'
                        )
                    }
                >
                    Trashed ({trashed.length})
                </Tab>
            </Tab.List>
            <Tab.Panels className="mt-2">

                <Tab.Panel
                    className={classNames(
                        'rounded-xl bg-white p-3',
                        'ring-white ring-opacity-60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2'
                    )}
                >
                    <Published data={published} update={update} setUpdate={setUpdate} />
                </Tab.Panel>
                <Tab.Panel
                    className={classNames(
                        'rounded-xl bg-white p-3',
                        'ring-white ring-opacity-60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2'
                    )}
                >
                    <Draft data={draft} update={update} setUpdate={setUpdate} />
                </Tab.Panel>
                <Tab.Panel
                    className={classNames(
                        'rounded-xl bg-white p-3',
                        'ring-white ring-opacity-60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2'
                    )}
                >
                    <Trashed data={trashed} update={update} setUpdate={setUpdate} />
                </Tab.Panel>
            </Tab.Panels>
        </Tab.Group>
    </div></Layout >
}

export default Post;