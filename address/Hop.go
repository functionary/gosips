/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : Hop.go
 * Author        : Rain Liu
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */
 
package address

import (

)

/**
 * The Hop interface defines a location a request can transit on the way to 
 * its destination, i.e. a route. It defines the host, port and transport of 
 * the location. This interface is used to identify locations in the 
 * {@link Router} interface.
 *
 * @see Router
 *
 * @author Sun Microsystems
 * @since 1.1
 *
 */

type Hop interface{

     /**
     * Returns the host part of this Hop.
     *
     * @return  the string value of the host.
     */     
    GetHost() string;

    /**
     * Returns the port part of this Hop.
     *
     * @return  the integer value of the port.
     */    
    GetPort() int;    

    /**
     * Returns the transport part of this Hop.
     *
     * @return the string value of the transport.
     */
    GetTransport() string;

    /**
     * This method returns the Hop as a string. 
     *
     * @return the stringified version of the Hop
     */    
    ToString() string;
    
}

