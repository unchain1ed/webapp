import React, { useState } from "react";
import { useRouter } from "next/router";
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
} from "@mui/material";
import axios from "axios";

const DeleteDialog = ({ id }) => {
  const router = useRouter();
  const [open, setOpen] = useState(false);

  const handleOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleDelete = async () => {
    try {
      // ブログ記事を削除するAPIエンドポイントにリクエストを送信
    //   await axios.delete(`http://localhost:8080/blog/delete/${blogId}`, {
    //   headers: {
    //     "Content-Type": "application/json", // JSON形式で送信するためのヘッダー設定
    //   },
    //     withCredentials: true,
    //   });

      const resoponse =await axios.get(`http://localhost:8080/blog/delete/${id}`, {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      withCredentials: true,
    });

    
      // 削除が成功した場合はブログ一覧ページにリダイレクト
      router.push("/blog/overview");
    } catch (error) {
      console.error("Error deleting blog:", error);
      // エラーハンドリングを行う場合はここに追加の処理を記述する
    }
  };

  return (
    <>
      <Button onClick={handleOpen}>Delete</Button>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>
          Are you sure you want to delete the blog post?(ブログ記事を削除しますか？)
        </DialogTitle>
        <DialogContent>
          <DialogContentText>
            Deleting the blog post is irreversible. Are you sure you want to proceed with the
            deletion?(ブログ記事を削除すると、元に戻すことはできません。本当に削除しますか？)
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={handleDelete} variant="contained" color="error">
            Delete
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default DeleteDialog;
