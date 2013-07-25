package header

import (

)

/**
 * This interface represents the Accept-Encoding request-header.
 * A client includes an AcceptEncodingHeader in a Request to tell the server
 * what coding schemes are acceptable in the Response e.g. compress, gzip.
 * <p>
 * If an AcceptEncodingHeader is present, and if the server cannot send a
 * Response which is acceptable according to the AcceptEncodingHeader, then
 * the server should return a Response with a status code of NOT_ACCEPTABLE.
 * <p>
 * An empty Accept-Encoding header field is permissible, it is equivalent to
 * <code>Accept-Encoding: identity</code>, meaning no encoding is permissible.
 * <p>
 * If no Accept-Encoding header field is present, the server SHOULD assume a
 * default value of identity.
 * <p>
 * For Example:<br>
 * <code>Accept-Encoding: gzip</code>
 *
 */
type AcceptEncodingHeader interface{
	ParametersHeader
	Encoding
	Header
   
    /**
     * Gets q-value of the encoding in this encoding value. A value of
     * <code>-1</code> indicates the<code>q-value</code> is not set.
     *
     * @return q-value of encoding value, -1 if q-value is not set.
     * @since v1.1
     */
    GetQValue() float32;

    /**
     * Sets q-value for the encoding in this encoding value. Q-values allow the
     * user to indicate the relative degree of preference for that encoding,
     * using the qvalue scale from 0 to 1. If no q-value is present, the
     * encoding should be treated as having a q-value of 1.
     *
     * @param qValue - the new float value of the q-value, a value of -1 resets
     * the qValue.
     * @throws InvalidArgumentException if the q parameter value is not
     * <code>-1</code> or between <code>0 and 1</code>.
     * @since v1.1
     */
    SetQValue( qValue float32) (InvalidArgumentException error);
    
    /**
     * Name of AcceptEncodingHeader
     */
    //public final static String NAME = "Accept-Encoding";

}