import * as React from 'react';
import Grid from '@mui/material/Grid';
import Typography from '@mui/material/Typography';
import Divider from '@mui/material/Divider';
import Markdown from './markdown';

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
}

export default function Main(props: MainProps) {
  const { posts } = props;
  return (
    <Grid
      item
      xs={12}
      md={8}
      sx={{
        '& .markdown': {
          py: 3,
        },
      }}
    >
      <Typography variant="h6" gutterBottom>
        {posts.length > 0 ? posts[0].title : ""}
      </Typography>
      <Divider />           
        <Markdown className="markdown" key={posts.length > 0 ? posts[0].title : ""}>
          {posts.length > 0 ? posts[0].content : ""}
        </Markdown>
    </Grid>
  );
}
