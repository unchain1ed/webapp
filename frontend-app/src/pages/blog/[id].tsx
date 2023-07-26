import React, { useEffect, useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";
import { Box, Typography } from "@mui/material";
import Chip from "@material-ui/core/Chip";
import Avatar from "@material-ui/core/Avatar";
import { useRouter } from "next/router";
import { GetServerSideProps } from "next";
import axios from "axios";
import format from "date-fns/format";
import NextLink from "next/link";
import { Logo } from "../../components/logo";

type Blog = {
  ID: string;
  LoginID: string;
  Title: string;
  Content: string;
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt: Date;
};

type BlogProps = {
  blog: Blog;
};

const useStyles = makeStyles((theme) => ({
  name: {
    lineHeight: 1,
  },
  content: {
    [theme.breakpoints.up("md")]: {
      paddingLeft: theme.spacing(8),
      paddingRight: theme.spacing(8),
    },
  },
  paragraph: {
    marginBottom: theme.spacing(3),
  },
  image: {
    maxWidth: "100%",
    borderRadius: theme.shape.borderRadius,
  },
}));

export default function Blog({ id }) {
  const classes = useStyles();
  const router = useRouter();
  // const { id } = router.query; // idを取得
  const [propsBlog, setBlogProps] = useState<Blog>({
    ID: "",
    LoginID: "",
    Title: "",
    Content: "",
    CreatedAt: new Date(),
    UpdatedAt: new Date(),
    DeletedAt: new Date(),
  });

  const content = {
    date: `${format(new Date(propsBlog.CreatedAt), "yyyy/MM/dd")}`,
    "header-p1": `${propsBlog.Title}`,
    // 'avatar': 'jpg',　//TODO
    name: `${propsBlog.LoginID}`,
    paragraph1: `${propsBlog.Content}`,
  };

  useEffect(() => {
    console.log(id)
    const getBlogContent = async () => {
      const hostname = process.env.NODE_ENV === "production" ? "server-app" : "localhost";

      try {
        const response = await axios.get(`http://${hostname}:8080/blog/overview/post/${id}`, {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          withCredentials: true,
        });
        
        const blog: Blog = response.data.blog;
        console.log(response.data.blog)
        console.log(blog)
        setBlogProps(blog)
        console.log(propsBlog)

      } catch (error) {
        console.error("エラーが発生しました", error);
        if (error.response.status === 302) {
          router.push("/auth/login");
        }
      }
    };
    getBlogContent();
  }, []);

  return (
    <section>
      <Box sx={{ p: 3 }}>
        <Box
          component={NextLink}
          href="/blog/overview"
          sx={{
            display: "inline-flex",
            height: 32,
            width: 32,
          }}
        >
          <Logo />
        </Box>
      </Box>
      <Container maxWidth="md">
        <Box py={10}>
          <Box textAlign="center" mb={5}>
            <Container maxWidth="sm">
              <Chip color="primary" label={content["date"]} />
              <Box my={4}>
                <Typography variant="h3" component="h2">
                  <Typography variant="h3" component="span" color="primary">
                    {content["header-p1"]}{" "}
                  </Typography>
                  <Typography variant="h3" component="span">
                    {content["header-p2"]}
                  </Typography>
                </Typography>
              </Box>
              <Box display="flex" justifyContent="center" alignItems="center">
                {/* <Avatar alt="" src={content['avatar']} /> */}
                <Box ml={2} textAlign="left">
                  <Typography variant="subtitle1" component="h2" className={classes.name}>
                    {content["name"]}
                  </Typography>
                  <Typography variant="subtitle1" component="h3" color="textSecondary">
                    {content["job"]}
                  </Typography>
                </Box>
              </Box>
            </Container>
          </Box>
          <Box className={classes.content}>
            <Typography
              variant="subtitle1"
              style={{ whiteSpace: "pre-wrap" }}
              color="textPrimary"
              paragraph={true}
              className={classes.paragraph}
            >
              {content["paragraph1"]}
            </Typography>
            <Box my={4}>
              <img src={content["image"]} alt="" className={classes.image} />
            </Box>
          </Box>
        </Box>
      </Container>
    </section>
  );
}
  export async function getServerSideProps(context) {
    // ここで context.params を使用して id を取得
    const { id } = context.params;

    // 例えば、この id を使ってデータを取得する処理などを行う

    // ページコンポーネントに props として id を渡す
    return {
      props: {
        id,
      },
    };
  }

// export const getServerSideProps: GetServerSideProps<BlogProps> = async (context) => {
//   const { id } = context.params; // idを取得
//   const hostname = process.env.NODE_ENV === "production" ? "server-app" : "localhost";

//   const response = await axios.get(`http://${hostname}:8080/blog/overview/post/${id}`, {
//     headers: {
//       "Content-Type": "application/x-www-form-urlencoded",
//     },
//     withCredentials: true,
//   });

//   const blog: Blog = response.data.blog;

//   return {
//     props: {
//       blog,
//     },
//   };
// };
