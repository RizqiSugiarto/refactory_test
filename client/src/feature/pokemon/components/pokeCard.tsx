import React from "react";
import { IPoke } from "../../../model/pokemon";

interface PokemonCardProps {
    pokemon: IPoke
}

const PokemonCard: React.FC<PokemonCardProps> = ({pokemon}) => {
    return(
        <div className="bg-white shadow-xl rounded-md p-3">
            <img className="w-64 h-64 rounded" src={pokemon.url} alt={pokemon.name} />
            <h2 className="text-xl font-bold text-center sm:text-left">{pokemon.name}</h2>
        </div>
    );
};

export default PokemonCard