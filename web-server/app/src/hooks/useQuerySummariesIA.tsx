import { useQuery } from "react-query"
import Config from "../config"
import { DataString } from "../model"
import { useCallback } from "react"
const toJson = (res: Response) => res.json()
const querySummarizeIA = (id: string) => () => fetch(`${Config.summarizeIAUri}/${id}`, {credentials: 'include'}).then(toJson)
// TODO: generic querys
export default (id: string) => {
    const query = useCallback(querySummarizeIA(id), [id])

    const { data } = useQuery<DataString>(
        [`${id}-ia`],
        query,
        {staleTime: 5*60*1000}, 
    )
    return data
}