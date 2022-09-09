import { Box } from "@mui/system";
import React from "react";
import TitleCard from "../DetailsTitleCard";

export function PodResourceView(props: { namespace: string, podName: string }) {
    const { podName, namespace } = props
    return <Box sx={{flexGrow: 1, p: 1}}>
        <TitleCard title={podName} />
    </Box>
}