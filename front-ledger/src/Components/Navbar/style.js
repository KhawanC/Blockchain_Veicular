import styled from 'styled-components'

export const Mainbox = styled.div`
    display: flex;

    width: 100vw;
    height: 50px;
`

export const LogoBox = styled.div`
    display: flex;

    width: 20%;
    height: 100%;
    padding-left: 10px;
`

export const LinksBox = styled.div`
    display: flex;
    width: 80%;
    height: 100%;

    align-items: center;
    justify-content: center;

    ul {
        display: flex;
        gap: 30px;
        padding-right: 35em;
    }

    ul li span {
        font-size: 17px;
    }

    ul li span:hover {
        color: white;
        cursor: pointer;
    }

`