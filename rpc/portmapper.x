/************************************/

const MNTPATHLEN = 1024;  /* Maximum bytes in a path name */
const MNTNAMLEN  = 255;   /* Maximum bytes in a name */
const FHSIZE3    = 64;    /* Maximum bytes in a V3 file handle */

typedef unsigned int youint;

typedef opaque fhandle3<FHSIZE3>;
typedef string dirpath<>;
typedef string name<MNTNAMLEN>;
/*

typedef struct mountbody *mountlist;

*/
struct mountbody {
	name       ml_hostname;
	dirpath    ml_directory;
	/* mountlist  ml_next; */
};


const PMAP_PORT = 111;

struct mapping {
    unsigned int prog;
    unsigned int vers;
    unsigned int prot;
    unsigned int port;
    bool mybool;
    float myfloat;
    double mydouble;
    /* struct mapping *pm; */
};

/*

const IPPROTO_TCP = 6;
const IPPROTO_UDP = 17;

struct pmaplist {
    mapping map;
    pmaplist *next;
};

struct pmaplist_first {
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

program PMAP {
    version V2 {
       void NULL(void)                  = 0;
       bool SET(mapping)                = 1;
       bool UNSET(mapping)              = 2;
       unsigned int GETPORT(mapping)    = 3;
       pmaplist_first DUMP(void)        = 4;
       call_result CALLIT(call_args)    = 5;
    } = 2;
} = 100000;
*/
