import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import axios from 'axios'
import styles from '../styles/Home.module.css'
import { useEffect } from 'react'


const Home: NextPage = () => {

  process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";
  // return (
  //   <div className={styles.container}>
  //     <Head>
  //       <title>Create Next App</title>
  //       <h1>NextPage</h1>
  //     </Head>
  //    </div>
  // )



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

  //以下のconsole.logはブラウザで実行されないサーバーサイドからの処理
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
