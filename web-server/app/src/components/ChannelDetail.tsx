import Typography from "@mui/material/Typography"
import useQuerySummaries from "../hooks/useQuerySummaries"
import { Box, Button, Card, CardContent, CardHeader, CircularProgress } from "@mui/material"
import useQuerySummariesIA from "../hooks/useQuerySummariesIA"
import { useState } from "react"

const styleCenter = {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
}
const style = {
    ...styleCenter,
    width: "100%",
    height: "100%",
}

const SumIA = ({ id }: { id: string | undefined}) => {
    if (!id){
        return null
    }
    const res = useQuerySummariesIA(id)

    if (!res) {
        return <Box sx={style}>
            <CircularProgress size="5rem"/>
            </Box>
    }
    return <Typography style={{whiteSpace: "pre-line"}} variant="body1" gutterBottom>
    {res.data}
    </Typography>
}

export default ({ id }: { id: string | undefined}) => {
    const [generate, setGenerate] = useState(false)
    
    if (!id){
        return null
    }

    const res = useQuerySummaries(id)

    if (!res) {
        return <Box sx={style}>
            <CircularProgress size="5rem"/>
            </Box>
    }

    
    return (
        <>
        <Card>
            <CardHeader  style={{ textAlign: 'center' }} title="Channel Resume:"/>
            <CardContent style={{ textAlign: 'center' }}>
            <Typography style={{whiteSpace: "pre-line"}} variant="body1" gutterBottom>
            {res.data}
            </Typography>
            </CardContent>
        </Card>   
        <Card>
            <CardHeader  style={{ textAlign: 'center' }} title="IA Resume:"/>
            <CardContent style={{ textAlign: 'center' }}>
            {
                !generate ?  <Button onClick={() => setGenerate(true)}>
                Generate
            </Button> :
            <SumIA id={id}/>
            } 
            </CardContent>
        </Card>   
        </>
 
    )
}