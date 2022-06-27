import React from 'react';

export const PostJs = ({posts, loading}) => {
    if(loading) {
        return <h2>Loading...</h2>
    }

    return(
        <>  <div className='p-36'>
                <table className='w-full text-left border-2 border-solid border-black'>
                    <thead className='text-xs text-gray-700 uppercase bg-blue-800'>
                        <tr>
                            <th className='px-6 py-3 text-2xl text-black'>Fabricante</th>
                            <th className='px-6 py-3 text-2xl text-black'>Co2 Emitido</th>
                            <th className='px-6 py-3 text-2xl text-black'>Hash</th>
                            <th className='px-6 py-3 text-2xl text-black'>VIM</th>
                        </tr>
                    </thead>
                    <tbody>
                        {posts.map(post => (
                            <tr key={post.id} className='bg-blue-600 border-2 border-solid border-black'>
                                <th scope='row' className='px-6 py-4 text-xl font-thin whitespace-nowrap'>
                                    {post.title}
                                </th>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </>
    )
};
