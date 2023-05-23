import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import axios from 'axios'
import styles from '../styles/Home.module.css'
import { useEffect } from 'react'


const Home: NextPage = () => {

  //TLS/SSL接続時に証明書を検証せずに接続を許可するかどうかを制御 "0" 検証を無効
  process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";

  useEffect(() => {
    
    (async() => {
    const item = await axios.get('https://localhost:8080/')
    })()
  } ,[])
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div>
        <h1>Hello World!</h1>
      </div>
    </main>
  )
}

export const getServerSideProps = async () => {

    //TLS/SSL接続時に証明書を検証せずに接続を許可するかどうかを制御 "0" 検証を無効
  process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";

    console.log("hello HH")
    const item = axios.get('https://localhost:8080/')
    return {
      props: {
        item: "hello world",
      },
    }
  }

export default Home
