import React from 'react';
import { Mainbox } from './style';

export const PaginacaoJs = ({postsPerPage, totalPosts, paginate}) => {
    const pageNumber = []

    for(let i = 1; i <= Math.ceil(totalPosts / postsPerPage); i++) {
        pageNumber.push(i)
    }
    return(
        <Mainbox>
            <ul>
                {pageNumber.map(num => (
                    <li key={num} onClick={() => paginate(num)} className='bg-blue-500'>
                        <span onClick={() => paginate(num)}>
                            {num}
                        </span>
                    </li>
                ))}
            </ul>
        </Mainbox>
    )
};
