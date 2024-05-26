import React from "react";
import useFetchPokemon from "./pokeApi";
import PokemonCard from "./components/pokeCard";

const PokePage: React.FC = () => {
    const {pokemon, error, loading} = useFetchPokemon()

    if(loading) return <div>Loading....</div>
    if(error) return <div>Error: {error.message}</div>

    return(
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
            {pokemon.map((poke, index) => (
                <PokemonCard key={index} pokemon={poke}/>
            ))}
        </div>
    )
}

export default PokePage