import styled from 'styled-components'

export const Mainbox = styled.div`
    background-image: linear-gradient(var(--primary-color), var(--secondary-color));

    display: flex;
    width: 100vw;
    height: 4.7em;

    .linha1 {
        height: 100%;
        width: 100%;

        display: flex;

        justify-content: space-between;
        align-items: center;
    }

    .linha1 .logo {

        display: flex;
        flex-direction: row;

        justify-content: center;
        align-items: center;

        margin: 1.5em;
        gap: 1.5em;
    }

    .linha1 .logo p {
        font-size: 1.7rem;
        font-weight: 600;
    }

    .linha1 .logo img {
        width: 30px;
        height: 30px;
    }

    .linha1 .links {
        display: flex;

        width: 40em;

        justify-content: space-between;
    }

    .linha1 .links ul {
        display: flex;

        width: 100%;
        
        flex-direction: row;
        justify-content: space-around;
    }

    .linha1 .links ul li{
        border: 1px solid black;
        border-radius: 25px;

        padding: 0.5rem;

        list-style-type: none;
        font-size: 1.5rem;
        font-weight: 500;
    }

`