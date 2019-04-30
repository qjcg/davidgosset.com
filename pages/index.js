import Head from 'next/head';
import Link from 'next/link';

function Home() {
  return (
    <div>
      <Head>
        <title>David Gosset</title>

        <meta charSet="utf-8" />
        <meta name="viewport" content="initial-scale=1.0, width=device-width" key="viewport" />

        <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Buenard" />
        <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Rufina" />

        <link rel="stylesheet" href="/static/main.css" />
      </Head>

      <h1>Welcome to davidgosset.com!</h1>

      <Link href="/teaching">
        <a>Teaching</a>
      </Link>

      <img src="/static/dave.jpg" alt="David Gosset" />
    </div>
  );
}

export default Home;
