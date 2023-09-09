import * as React from "react";
import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Markdown from "./markdown";
import { useEffect, useState } from "react";
import closestIndexTo from "date-fns/esm/closestIndexTo";

type Blog = {
  id: string;
  loginID: string;
  title: string;
  content: string;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date;
};

interface MainProps {
  posts: Blog[];
  id: any;
}

export default function Main(props: MainProps) {
  const { posts } = props;
  const { id } = props;

  // // id に応じて表示内容を選択
  // let selectedPost = id && posts[id] ? posts[id] : (posts.length > 0 ? posts[0] : null);

  const [selectedPost, setSelectedPost] = useState<Blog | null>(null);

  // idが変更されたときに選択された投稿を更新
  useEffect(() => {
    if (id && posts.length > 0) {
     

      // findIndexメソッドを使用してidを検索し、indexを取得
      const index = posts.findIndex((blog) => blog.id === id);


      setSelectedPost(posts[index]);
    } else if (posts.length > 0) {
      setSelectedPost(posts[0]);
    } else {
      // const newPost = Blog[0].title = "aaa";
      // setSelectedPost(newPost);
    }
  }, [id, posts]);

  if (!selectedPost) {
    // 選択された投稿が存在しない場合の処理
    return <div>No posts available</div>;
  }

  return (
    <Grid
      item
      xs={12}
      md={8}
      sx={{
        "& .markdown": {
          py: 3,
        },
      }}
    >
      <Typography variant="h6" gutterBottom>
        {selectedPost.title}
      </Typography>
      <Divider />
      <Markdown className="markdown" key={selectedPost.title}>
        {selectedPost.content}
      </Markdown>
    </Grid>
  );
}

// export default function Main(props: MainProps, id: string) {
//   const { posts } = props;
//   return (
//     <Grid
//       item
//       xs={12}
//       md={8}
//       sx={{
//         '& .markdown': {
//           py: 3,
//         },
//       }}
//     >
//       <Typography variant="h6" gutterBottom>
//         {posts.length > 0 ? posts[0].title : ""}
//       </Typography>
//       <Divider />
//         <Markdown className="markdown" key={posts.length > 0 ? posts[0].title : ""}>
//           {posts.length > 0 ? posts[0].content : ""}
//         </Markdown>
//     </Grid>
//   );
// }
