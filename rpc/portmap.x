/*
 * From RFC1833
 */

const PMAP_PORT = 111;      /* portmapper port number */
/* const FOOBAR = TRUE; */

struct mapping {
       /* int foobar; */
       unsigned int prog;
       unsigned int vers;
       unsigned int prot;
       unsigned int port;
};

const IPPROTO_TCP = 6;      /* protocol number for TCP/IP */
const IPPROTO_UDP = 17;     /* protocol number for UDP/IP */

struct pmaplist {
    mapping map;
    pmaplist *next;
};

struct dump_result {
    pmaplist *next;
};

struct call_args {
       unsigned int prog;
       unsigned int vers;
       unsigned int proc;
       opaque args<>;
};

struct call_result {
    unsigned int port;
    opaque res<>;
};

/*
program PMAP {
     version V2 {
        void NULL(void)               = 0;
        bool SET(mapping)             = 1;
        bool UNSET(mapping)           = 2;
        unsigned int GETPORT(mapping) = 3;
        dump_result DUMP(void)        = 4;
        call_result CALLIT(call_args) = 5;
     } = 2;
} = 100000;
*/
