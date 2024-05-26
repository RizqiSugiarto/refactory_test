import axios from "axios"
import { useState, useEffect } from "react"
import { IPoke } from "../../model/pokemon"

interface FetchPokemonReturn {
    pokemon: IPoke[]
    loading: boolean
    error: Error | null
}

const useFetchPokemon = (): FetchPokemonReturn => {
    const[pokemon, setPokemon] = useState<IPoke[]>([])
    const[loading, setLoading] = useState(true)
    const[error, setError] = useState<Error | null>(null)

    useEffect(() => {
        const fetchPokemon = async () => {
            try {
                const response = await axios.get<IPoke[]>("http://localhost:8080/pokemon")
                setPokemon(response.data)
            } catch (error) {
                setError(error as Error)
            }finally{
                setLoading(false)
            }
        }

        fetchPokemon()
    }, [])
    return {pokemon, loading, error}
}

export default useFetchPokemon