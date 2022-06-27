import styled from 'styled-components'

export const Mainbox = styled.div`
    background-image: linear-gradient(var(--primary-color), var(--secondary-color));
    
    display: flex;
    flex-direction: column;

    justify-content: center;
    align-items: center;

    width: 35em;
    height: 12em;

    .linha1 p{
        text-align: center;
        font-size: 1.9rem;
        font-weight: 600;
    }
`;