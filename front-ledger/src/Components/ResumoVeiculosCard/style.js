import styled from 'styled-components'

export const Mainbox = styled.div`
    background-image: linear-gradient(var(--primary-color), var(--secondary-color));
    
    display: flex;

    justify-content: space-around;
    align-items: center;

    border: 1px solid white;
    border-radius: 30px;

    width: 35em;
    height: 15em;

    .linha1 p{
        font-size: 2.1rem;
        font-weight: 600;
    }
`;