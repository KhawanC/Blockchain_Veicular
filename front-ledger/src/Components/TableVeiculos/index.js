import React, { useState, useEffect } from 'react';
import axios from 'axios'
import { PostJs } from '../PostJs';
import { PaginacaoJs } from '../PaginacaoJS';

export const TableVeiculos = (params) => {
    const [posts, setPosts] = useState([]);
    const [loading, setLoading] = useState(false);
    const [currentPage, setCurrentPage] = useState(1);
    const [postsPerPage] = useState(10);

    useEffect(() => {
        const fetchPost = async () => {
            setLoading(true);
            const res = await axios.get('https://jsonplaceholder.typicode.com/posts');
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
            <div className='h-screen bg-blue-800'>
                <PostJs posts={currentPost} loading={loading}/>
                <PaginacaoJs postsPerPage={postsPerPage} totalPosts={posts.length} paginate={paginate}/>
            </div>   
        </>
    )
};
