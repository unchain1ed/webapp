import { useCallback, useEffect, useState } from 'react';
import {
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  Divider,
  TextField,
  Unstable_Grid2 as Grid
} from '@mui/material';
import axios from 'axios';
import { useRouter } from "next/router";

type User = {
  ChangeId: string;
  NowId: string;
};

export const AccountProfileDetails = () => {
  const router = useRouter();
  const [userId, setUserId] = useState("");
  const [values, setValues] = useState({
    id: userId
  });
  const [propsUser, setUserProps] = useState<User>({
    ChangeId: "changeId",
    NowId: "nowId"
  });
  const [isPosting, setPosting] = useState(false);
  //user.ChangeId, user.NowId

  // const handleChange = useCallback(
  //   (event) => {
  //     setValues((prevState) => ({
  //       ...prevState,
  //       [event.target.name]: event.target.value
  //     }));
  //   },
  //   []
  // );

  const handleSubmit = useCallback(
    (event) => {
      event.preventDefault();
    },
    []
  );

  useEffect(() => {
    const fetchData = async () => {
      const hostname = process.env.NODE_ENV === "production" ? "server-app" : "localhost";
      try {
        const response = await axios.get(`http://${hostname}:8080/api/login-id`, {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          withCredentials: true,
        });
        const reqNowId = response.data.id;
        // setUserProps(nowId);
        setUserProps({
          ...propsUser, // 既存のuserIdオブジェクトを展開して新しいオブジェクトを作成
          NowId: reqNowId,
          ChangeId: reqNowId // 新しい値をidフィールドに代入
        });
      } catch (error) {
        console.error(error);
      }
    };
    // コンポーネントのマウント時にリクエストを実行
    fetchData();
  }, []);

  const handlePost = async () => {
    if (isPosting) {
      // 追加: 処理が実行中の場合は何もしない
      return;
    }
    setPosting(true); // 追加: 処理を実行中にセット

    const hostname = process.env.NODE_ENV === "production" ? "server-app" : "localhost";
    try {
      const response = await axios.post(
        `http://${hostname}:8080/update/id`,
        propsUser, 
        {
          headers: {
            "Content-Type": "application/json", // JSON形式で送信するためのヘッダー設定application/json"でないと飛ばない
          },
          withCredentials: true,
        }
      );
      console.log("IDを変更しました");
    } catch (error) {
      // エラー処理を行う場合はここに記述
      console.error("Error editing id:", error);
    } finally {
      // 処理が終了したので false にセット
      setPosting(false);
      router.push("/");
    }
  };

  return (
    <form
      autoComplete="off"
      noValidate
      onSubmit={handlePost}
    >
      <Card>
        <CardHeader
          subheader="ID can be edited"
          title="Profile"
        />
        <CardContent sx={{ pt: 0 }}>
          <Box sx={{ m: -1.5 }}>
            <Grid
              container
              spacing={3}
            >
              <Grid
                xs={12}
                md={6}
              >
                <TextField
                  fullWidth
                  label="ID"
                  name="id"
                  onChange={(event) => {
                    setUserProps({
                      ...propsUser, // 既存のuserIdオブジェクトを展開して新しいオブジェクトを作成
                      ChangeId: event.target.value, // 新しい値をidフィールドに代入
                    });
                  }}
                  required
                  value={propsUser.ChangeId}
                />
              </Grid>
              <Grid
                xs={12}
                md={6}
              >
              </Grid>
            </Grid>
          </Box>
        </CardContent>
        <Divider />
        <CardActions sx={{ justifyContent: 'flex-end' }}>
          <Button 
          variant="contained"
          onClick={handlePost}
          disabled={!propsUser}
          >
            Save details
          </Button>
        </CardActions>
      </Card>
    </form>
  );
};
