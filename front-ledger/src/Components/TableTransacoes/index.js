import React, { useState, useEffect } from 'react';
import { Api } from '../../Services/Api';
import { PostTransJs } from '../PostTransJs';
import { PaginacaoJs } from '../PaginacaoJS';
import { Mainbox } from './style';

export const TableTransacoes = (params) => {
    const [posts, setPosts] = useState([]);
    const [loading, setLoading] = useState(false);
    const [currentPage, setCurrentPage] = useState(1);
    const [postsPerPage] = useState(10);

    useEffect(() => {
        const fetchPost = async () => {
            setLoading(true);
            const res = await Api.get('ordem');
            console.log(res)
            setPosts(res.data);
            setLoading(false);
        }

        fetchPost();
    }, [])

    const paginate = (pageNumber) => setCurrentPage(pageNumber);

    const indexLastPost = currentPage * postsPerPage;
    const indexFirstPost = indexLastPost - postsPerPage
    const currentPost = posts.slice(indexFirstPost, indexLastPost)

    return(
        <>
            <Mainbox>
                <PostTransJs posts={currentPost} loading={loading}/>
                <PaginacaoJs postsPerPage={postsPerPage} totalPosts={posts.length} paginate={paginate}/>
            </Mainbox>   
        </>
    )
};
