package com.company.process;

import java.util.*;

/**
 * @author  Arturo Negreiros
 * @version Java 11
 */
public class Main {

    private final HashMap<String, Object> socialMediaA = new HashMap<>();
    private final HashMap<String, Object> socialMediaB = new HashMap<>();
    private final HashMap<String, Object> socialMediaC = new HashMap<>();
    private final ArrayList<HashMap> usersAndFollowers;



    public Main(){

        this.socialMediaA.put("User", "userA");
        this.socialMediaA.put("Following", new ArrayList<>(Arrays.asList("userB", "userD","userE", "userG")));
        this.socialMediaB.put("User", "userB");
        this.socialMediaB.put("Following", new ArrayList<>(Arrays.asList("userC", "userJ","userI", "userE")));
        this.socialMediaC.put("User", "userC");
        this.socialMediaC.put("Following", new ArrayList<>(Arrays.asList("userM", "userN","userJ", "userI", "userE")));

        this.usersAndFollowers = new ArrayList<>(Arrays.asList(this.socialMediaA, this.socialMediaB, this.socialMediaC));
    }


    /**
     *
     * @param userStart el usuario del cuadl deseamos saber la distancia al usuario destino
     * @param userDestiny usuario al que debemos encontrar su distancia con respecto al usuario inicial
     * @return un entero que será la distancia entre el usuario inicial y el usuario final
     */
    public int calculateFollowers(String userStart, String userDestiny){

        /*
         *
         */


        int counter = 0; // iniciamos una variable en 0

        /*
        * Luego dos arreglos de tipo entero y cadena de texto respectivamente
        * */
        ArrayList<Integer> roadMaps = new ArrayList<>();
        ArrayList<String> users = new ArrayList<>();

        // el bloque try catch en caso de una "Exception"
        try {

            // un For each para iterar cada elemento de la lista de usuarios y seguidores
            for(HashMap value: this.usersAndFollowers){

                /*verificamos qué usuario está siguiendo al usuario final del que queremos saber la distancia
                * luego lo agregamos a el arreglo "user" para guardar los usuarios que siguen al usuario final ejemplo:
                * del usuario A al usuario M -> usuario inicial A y usuario final M
                */
                if (((ArrayList<String>)value.get("Following")).contains(userDestiny)){
                    users.add((String) value.get("User"));
                }

            }
            // como queremos saber la distancia la primera referencia es que aumentará en uno la distancia
            counter++;

            /*verificamos si el usuario inicial está en el arreglo de los que están siguiendo al usuaruio final
            * si es así, entonces devolvemos el valor de la distancia
            * */
            if(users.contains(userStart)){
                return counter;
            }else{

                /*
                * creamos otra varibale de tipo entero en 0
                * solo para evitar que se acumule el contador inicial "counter"
                * */

                int counter2 = 0;
                // creamos una etiqueta que sea referenciada a los 'for´s' anidados y poder cerrarlos al mismo tiempo

                outerloop:
                for(String user: users){

                    // esta variable es para saber quien sigue al primer elemtno de la iteracion "for(String user: users)"
                    String whoFollowUser = "";
                    for (HashMap value: this.usersAndFollowers){

                        /* en caso de que el usuario inicial esté siguiendo a la primera iteración, entonces cerramos los dos
                        * los dos 'for´s' con un break general
                        *
                        *  */
                        if(((ArrayList<String>)value.get("Following")).contains(user)){
                            if(value.get("User").equals(userStart)){
                                counter2++;
                                break outerloop; // instancia para cerrar a los dos "for´s"
                            }else{
                                // aumentamos el contador
                                counter2++;

                                // instanciamos el valor para verificar en el siguiente bloque quien sigue a ese usuario
                                whoFollowUser = (String) value.get("User");
                            }
                        }
                    }
                    for(HashMap value: this.usersAndFollowers){

                        // verificamos que el usuario sea diferente del que ya está iterando
                        if (!value.get("User").equals(user)){

                            // verificamos que lo contenga en los "seguidos"
                            if(((ArrayList<String>)value.get("Following")).contains(whoFollowUser)){
                                if (value.get("User").equals(userStart)){
                                    // si es así aumentamos el contador, y cerramos el bloque
                                    counter2++;
                                    break outerloop;
                                }
                            }

                        }
                    }
                }
                // agregamos al roadmap la suma de contadores y luego las ordemaos con Collections.sort()
                roadMaps.add(counter + counter2);
            }
            Collections.sort(roadMaps);

            // retornamos el primer valor luego de que se haya ordenado, que seria el menor valor
            return roadMaps.get(0);
        }catch (Exception e){
            e.printStackTrace();
            return  0;
        }

    }

    public static void main(String[] args) {

        Main main = new Main();
        System.out.println(main.calculateFollowers("userA", "userM")); // 3
        System.out.println(main.calculateFollowers("userA", "userI")); // 2
        System.out.println(main.calculateFollowers("userA", "userN")); // 3


    }
}
